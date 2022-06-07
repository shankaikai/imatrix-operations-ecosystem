
timing = '["18:00:00-23:59:59", "00:00:00-6:00:00"]'
query = "INSERT INTO " + \
"operations_ecosystem.availability(week, year, guard, sunday, monday, tuesday, wednesday, thursday, friday, saturday, next_sunday) " + \
"VALUES({}, 2022, {},'{}', '{}','{}','{}','{}','{}','{}','{}');"

with open("avail.txt", "w") as f:
    for userId in range(1, 15):
        for week in range(1, 56):
            f.write(query.format(week, userId, timing, timing, timing, timing, timing, timing, timing, timing))
            f.write("\n")

