from logging.handlers import BaseRotatingHandler
import pandas as pd
import numpy as np
from math import exp
import datetime
import calendar
import plotly.express as px
import plotly.graph_objects as go

import dash
import dash_core_components as dcc
import dash_html_components as html
import dash_table as dt
from dash.dependencies import Input, Output

app = dash.Dash(__name__)

# get last week or last month
week = pd.to_datetime(datetime.datetime.now()).week - 1 # takes previous week
weeks = list(range(week-3,week+1))
month = pd.to_datetime(datetime.datetime.now()).month - 1 # takes previous month
months = list(range(month-3,month+1))

# ---------------------------------------------------------------------
# Import data
users = pd.read_csv("users.csv")
broadcast = pd.read_csv("broadcast_may_to_jul.csv")
broadcast_recipients = pd.read_csv("broadcast_recepients_may_to_jul.csv")
schedule_detail = pd.read_csv("schedule_detail_may_to_jul.csv")

# Change data types
broadcast["creation_date"] = pd.to_datetime(broadcast["creation_date"], dayfirst=True)
broadcast["deadline"] = pd.to_datetime(broadcast["deadline"], dayfirst=True)
broadcast_recipients["last_replied"] = pd.to_datetime(broadcast_recipients["last_replied"], dayfirst=True)
schedule_detail["custom_start_time"] = pd.to_datetime(schedule_detail["custom_start_time"], dayfirst=True)
schedule_detail["custom_end_time"] = pd.to_datetime(schedule_detail["custom_end_time"], dayfirst=True)
schedule_detail["attendance_time"] = pd.to_datetime(schedule_detail["attendance_time"], dayfirst=True)

# ---------------------------------------------------------------------
# Functions

def create_guard_list(users):
    results = list()
    for employee in range(len(users[(users["user_type"]=="I-Specialist")])):
        dct = {"label": users["name"][employee], "value": users["user_id"][employee]}
        results.append(dct)
    return results

def calc_attendance(guard_id):
    guard_schedule = attendance_table_filtered[(attendance_table_filtered["guards_assigned"]==guard_id)]
    count_punctual = guard_schedule[(guard_schedule["attendance_time"]<=guard_schedule["custom_start_time"])].count()["attended"]
    count_late = guard_schedule[(guard_schedule["attendance_time"]>guard_schedule["custom_start_time"])].count()["attended"]
    count_absent = guard_schedule[(guard_schedule["attended"]==0)].count()["guards_assigned"]
    return [count_punctual, count_late, count_absent]

def calc_avg_response_time(guard_id):
    guard_messages = broadcast_table_filtered[(broadcast_table_filtered["recipient"]==guard_id)]
    avg_response_time = int(guard_messages["response_time"].mean())
    return avg_response_time


# To transform tables
def calc_response_time(row):
    if row["acknowledged"]==1:
        response_time = (row["last_replied"]-row["creation_date"]).total_seconds()
    else:
        response_time = float("inf")
    return response_time

def calc_lateness(row):
    if row["confirmation"]==1 & row["attended"]==1:
        lateness = max(0, (row["attendance_time"]-row["custom_start_time"]).total_seconds()/60)
    elif row["confirmation"]==1 & row["attended"]==0:
        lateness = float("inf")
    else:
        lateness = None
    return lateness

def calc_hours_worked(row):
    if row["confirmation"]==1 & row["attended"]==1:
        hours_worked = (row["custom_end_time"]-max(row["custom_start_time"],row["attendance_time"])).total_seconds()/(3600)
    else:
        hours_worked = 0
    return hours_worked

def avg_hours_worked_weekly(): # avg of all guards, returns list of 4 weeks
    avg_hours_worked = []
    for i in weeks:
        hours = attendance_table_filtered[(attendance_table_filtered["week"]==i)].groupby("guards_assigned")["hours_worked"].sum().mean()
        avg_hours_worked.append(hours)
    return avg_hours_worked

def hours_worked_weekly(guard_id):
    hours_worked = []
    for i in weeks:
        hours = attendance_table_filtered[(attendance_table_filtered["guards_assigned"]==guard_id) & (attendance_table_filtered["week"]==i)]["hours_worked"].sum()
        hours_worked.append(hours)
    return hours_worked

def avg_hours_worked_monthly():
    avg_hours_worked = []
    for i in months:
        hours = attendance_table_filtered_month[(attendance_table_filtered_month["month"]==i)].groupby("guards_assigned")["hours_worked"].sum().mean()
        avg_hours_worked.append(hours)
    return avg_hours_worked

def hours_worked_monthly(guard_id):
    hours_worked = []
    for i in months:
        hours = attendance_table_filtered_month[(attendance_table_filtered_month["guards_assigned"]==guard_id) & (attendance_table_filtered_month["month"]==i)]["hours_worked"].sum()
        hours_worked.append(hours)
    return hours_worked
    

def calc_scores(row):
    value = row["value"]
    if row["score_type"]=="lateness":
        score = 5*exp(-0.02*value)
        return score
    elif row["score_type"]=="response_time":
        score = 5*exp(-0.002*value)
        return score
    elif row["score_type"]=="hours_worked":
        avg_hours_dict = dict(zip(weeks, avg_hours_worked_weekly()))
        wk = row["week"]
        avg = avg_hours_dict[wk]
        score = 5 - 5*abs((value-avg)/avg)
        return score

# ---------------------------------------------------------------------
# Tables needed for graph / values

# get only confirmed shifts
schedule_detail_confirmed = schedule_detail[(schedule_detail["confirmation"]==1)]

# add lateness and hours_worked
attendance_table = schedule_detail_confirmed.copy()
attendance_table["lateness"] = attendance_table.apply(lambda row: calc_lateness(row), axis=1)
attendance_table["hours_worked"] = attendance_table.apply(lambda row: calc_hours_worked(row), axis=1)

# get only broadcast messages with high urgency
broadcast_merge = pd.merge(broadcast[["broadcast_id", "creation_date", "type", "urgency"]], 
                            broadcast_recipients, 
                            left_on="broadcast_id", 
                            right_on="related_broadcast").drop("related_broadcast", 
                            axis=1)
broadcast_urgent = broadcast_merge[(broadcast_merge["urgency"]=="High")]

# add response_time
broadcast_table = broadcast_urgent.copy()
broadcast_table["response_time"] = broadcast_table.apply(lambda row: calc_response_time(row), axis=1)

# add day, month, week, year
broadcast_table["day"] = broadcast_table.apply(lambda row: row["creation_date"].day, axis=1)
broadcast_table["month"] = broadcast_table.apply(lambda row: row["creation_date"].month, axis=1)
broadcast_table["year"] = broadcast_table.apply(lambda row: row["creation_date"].year, axis=1)
broadcast_table["week"] = broadcast_table.apply(lambda row: row["creation_date"].week, axis=1)

attendance_table["day"] = attendance_table.apply(lambda row: row["custom_start_time"].day, axis=1)
attendance_table["month"] = attendance_table.apply(lambda row: row["custom_start_time"].month, axis=1)
attendance_table["year"] = attendance_table.apply(lambda row: row["custom_start_time"].year, axis=1)
attendance_table["week"] = attendance_table.apply(lambda row: row["custom_start_time"].week, axis=1)

# Filter for last 4 weeks
attendance_table_filtered = attendance_table[(attendance_table["week"]>week-4) & (attendance_table["week"]<=week)]
broadcast_table_filtered = broadcast_table[(broadcast_table["week"]>week-4) & (broadcast_table["week"]<=week)]

# Filter for last 3 months
attendance_table_filtered_month = attendance_table[(attendance_table["month"]>month-4) & (attendance_table["month"]<=month)]

# ---------------------------------------------------------------------
# Score calculations

guard_list = create_guard_list(users)
score_table_cols = ["guard_id", "score_type", "value", "week", "year"]

response_time_table = broadcast_table_filtered.loc[:, ["recipient", "response_time", "week", "year"]]
response_time_table.insert(1, "score_type", "response_time")
response_time_table.columns = score_table_cols

lateness_table = attendance_table_filtered.loc[:, ["guards_assigned", "lateness", "week", "year"]]
lateness_table.insert(1, "score_type", "lateness")
lateness_table.columns = score_table_cols

hours_worked_table = attendance_table_filtered.loc[:, ["guards_assigned", "hours_worked", "week", "year"]]
hours_worked_table.insert(1, "score_type", "hours_worked")
hours_worked_table.columns = score_table_cols
hours_worked_table = hours_worked_table.groupby(["guard_id", "score_type", "week", "year"]).sum().reset_index()
hours_worked_table = hours_worked_table[score_table_cols]

score_table = pd.concat([lateness_table, response_time_table, hours_worked_table])

avg_hours_dict = dict(zip(weeks, avg_hours_worked_weekly()))

score_table["score"] = score_table.apply(lambda row: calc_scores(row), axis=1)

final_score_table = score_table.loc[:, ["guard_id", "score_type", "score"]]
final_score_table = final_score_table.groupby(["guard_id", "score_type"]).mean().reset_index()

def calc_final_score(guard_id, df=final_score_table): # df = final_score_table
    df = df[(df["guard_id"]==guard_id)]
    lateness_score = df[(df["score_type"]=='lateness')]["score"].values[0]
    response_time_score = df[(df["score_type"]=='response_time')]["score"].values[0]
    hours_worked_score = df[(df["score_type"]=='hours_worked')]["score"].values[0]
    final_score = 0.4*lateness_score+0.3*response_time_score+0.3*hours_worked_score
    return round(final_score,2)

final_scores = pd.DataFrame(columns=["guard_id", "score"])
final_scores["guard_id"]=list(final_score_table["guard_id"].unique())
final_scores["score"] = final_scores.apply(lambda row: calc_final_score(row["guard_id"]), axis=1)

employees = users[(users["user_type"]=="I-Specialist")].loc[:, ["user_id", "name"]]
employee_score_table = pd.merge(employees, final_scores, left_on="user_id", right_on="guard_id").iloc[:,[1,3]]
employee_score_table.columns = ["I-Specialist", "Score"]

counts, bins = np.histogram(final_scores["score"], bins=range(0,6,1))
bins = 0.5 * (bins[:-1] + bins[1:])

# ---------------------------------------------------------------------
# App Layout
app.layout = html.Div([
    html.H1("iMatrix Employees Dashboard", style={'text-align': 'center'}),
    html.Div([
        html.Div([ 
            html.P("Select the I-Specialist:", style={'font-weight': 'bold'}),
            dcc.Dropdown(id="select_guard",
                options=guard_list,
                multi=False,
                value=1,
                )
            ], style={'width':'20%'}),

        html.Div([
            # score div
            html.Div([
                html.H3("Employee Score", style={'text-align': 'center'}),
                html.H2(id='final_score', children=[], style={'text-align': 'center'}),
                html.P(id="stars", children=[], style={'text-align': 'center'})
            ]),

            # response time div
            html.Div([
                html.H3("Average Response Time", style={'text-align': 'center'}),
                html.H2(id='avg_response_time', children=[], style={'text-align': 'center'}) # add distribution of response times?
            ], style={'margin-left':'20%'})  
        ], className="flex-row", style={'width':'60%', 'justify-content': 'center'})
                           
                ], className="flex-row"),
   


    # First row
    html.Div([

        html.Div([
            dcc.Graph(id='attendance', className="attendance-graph")
        ], className="attendance"),

        html.Div([
            dcc.RadioItems(id = 'radio_items',
                        labelStyle= {"display": "inline-block"},
                        value = 'Weekly',
                        options = [{'label': 'Weekly', 'value': 'Weekly'},
                                    {'label': 'Monthly', 'value': 'Monthly'}],
                        style = {'text-align': 'center', 'margin': '10px'}),

            dcc.Graph(id='working_hours', figure={}, style={'width': '90%'})
        ], className="working-hours"),
        

    ], className="flex-row"), 

    # Second row
    html.Div([
        html.Div(
            # Table of scores
            dt.DataTable(data=employee_score_table.to_dict('records'),
                        columns=[{"name": i, "id": i} for i in  employee_score_table.columns],
                        sort_action='native',
                        style_as_list_view = True,
                        style_header = {'textAlign': 'center', 'fontWeight': 'bold'},
                        style_cell={'textAlign': 'center', 'color': 'white','backgroundColor': '#111111'},
                        style_data_conditional=[
                            {
                                'if': {
                                    'filter_query': '{{Score}} = {}'.format(employee_score_table['Score'].min()),
                                },
                                'backgroundColor': '#FF4136',
                                'color': 'white'
                            },
                            {
                                'if': {
                                    'filter_query': '{{Score}} = {}'.format(employee_score_table['Score'].max()),
                                },
                                'backgroundColor': '#00C400', 
                                'color': 'white'
                            },

                        ], 
                        style_table={'width':'100%'}),
            className="datatable"
        ),

        html.Div(
            # Distribution of scores
            dcc.Graph(id='score_graph', figure={}, style={'width': '90%'}),
            className="scores"
        )

    ], className="flex-row")

], className="container")

# ---------------------------------------------------------------------
# Connect the Plotly graphs with Dash components
@app.callback(
    [Output(component_id='attendance', component_property='figure'),
    Output(component_id='avg_response_time', component_property='children'),
    Output(component_id='working_hours', component_property='figure'),
    Output(component_id='final_score', component_property='children'),
    Output(component_id='stars', component_property='children'),
    Output(component_id='score_graph', component_property='figure')],
    [Input(component_id='select_guard', component_property='value'),
    Input(component_id='radio_items', component_property='value')]
)

def update_graph(guard_selected, frequency):
    attendance_values = calc_attendance(guard_selected)
    dff = pd.DataFrame({'names' : ['Punctual','Late', 'Absent'], 'values' :  attendance_values})
    fig = px.pie(dff, values='values', names = 'names', hole=0.5, 
                color_discrete_sequence = ['green', 'orange','grey'], title="Attendance", labels = None, height=350, template= 'plotly_dark')
    fig.update_layout(title_x=0.5)
    
    response_time = "{} seconds".format(calc_avg_response_time(guard_selected))

    if frequency=="Weekly":
        fig2_x = ["Week "+ str(x) for x in weeks]
        fig2 = px.line(x=fig2_x, y = avg_hours_worked_weekly(), color=px.Constant("Average"),
                    labels=dict(x="Week of Year", y="Hours"), title="Hours worked in a week", height=330, template= 'plotly_dark')
        fig2.add_bar(x=fig2_x, y=hours_worked_weekly(guard_selected), name=guard_list[guard_selected-1]['label'])
    if frequency=="Monthly":
        fig2_x = [calendar.month_name[x] for x in months]
        fig2 = px.line(x=fig2_x, y = avg_hours_worked_monthly(), color=px.Constant("Average"),
                    labels=dict(x="Month", y="Hours"), title="Hours worked in a month", height=330, template= 'plotly_dark')
        fig2.add_bar(x=fig2_x, y=hours_worked_monthly(guard_selected), name=guard_list[guard_selected-1]['label'])
    
    fig2.update_layout(title_x=0.5)
    final_score = final_scores[(final_scores["guard_id"]==guard_selected)]["score"].item()
    final_score_guard = "{:.2f}".format(final_score)
    stars = int(final_score)*"⭐"

    score_graph = px.bar(x=bins, y = counts, labels={'x': 'Count', 'y': 'Score'}, title='Distribution of scores', color_discrete_sequence=["#f5bb38"], template= 'plotly_dark')
    score_graph.update_layout(title_x=0.5)
    
    return fig, response_time, fig2, final_score_guard, stars, score_graph 

# ---------------------------------------------------------------------
if __name__ == '__main__':
    # app.run_server(debug=True)
    app.run_server()