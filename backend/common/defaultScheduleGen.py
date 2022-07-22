start_timing = '1965-01-{} 18:00:00'
end_timing = '1965-01-{} 06:00:00'
schedule_query = "INSERT INTO " + \
"operations_ecosystem.schedule(aifs_id, start_time, end_time) " + \
"VALUES({}, '{}','{}');"

schedule_detail_query = "INSERT INTO " + \
"operations_ecosystem.schedule_detail(schedule, guard_assigned, custom_start_time, custom_end_time, is_assigned, rejected) " + \
"VALUES({}, '{}','{}','{}', 1, 0);"

aifs_client_schedule_query = "INSERT INTO " + \
"operations_ecosystem_testing.aifs_client_schedule(schedule, related_client, patrol_order) " + \
"VALUES({}, '{}', {});"

default_rostering_query = "INSERT INTO " + \
"operations_ecosystem.default_rostering(day_of_week, aifs1_schedule, aifs2_schedule, aifs3_schedule) " + \
"VALUES({}, '{}', '{}', '{}');"

count = 0
with open("default_schedule.txt", "w") as f:
    for day in range(0, 7):
        for aifs in range(1, 4):
            count +=1
            # f.write(schedule_query.format(aifs, start_timing.format(3+day), end_timing.format(4+day)))
            # f.write("\n")
            # # guard ids start from 3
            # f.write(schedule_detail_query.format(count, aifs+2, start_timing.format(3+day), end_timing.format(4+day)))
            # f.write("\n")
            # Client and AIFS have the same id
            # have different patrol order for different aifs
            if aifs == 1:
                f.write(aifs_client_schedule_query.format(count, 1, 1))
                f.write("\n")
                f.write(aifs_client_schedule_query.format(count, 2, 2))
                f.write("\n")
                f.write(aifs_client_schedule_query.format(count, 3, 3))
                f.write("\n")            
            if aifs == 2:
                f.write(aifs_client_schedule_query.format(count, 1, 2))
                f.write("\n")
                f.write(aifs_client_schedule_query.format(count, 2, 3))
                f.write("\n")
                f.write(aifs_client_schedule_query.format(count, 3, 1))
                f.write("\n")            
            if aifs == 3:
                f.write(aifs_client_schedule_query.format(count, 1, 3))
                f.write("\n")
                f.write(aifs_client_schedule_query.format(count, 2, 1))
                f.write("\n")
                f.write(aifs_client_schedule_query.format(count, 3, 2))
                f.write("\n")
        # f.write(default_rostering_query.format(day, count-2, count-1 , count))
        # f.write("\n")