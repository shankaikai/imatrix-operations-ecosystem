// Get the dates from the tuesday to sunday of the week and put them in an array
export default function getRosterDates(offset: number): Date[] {
  const today = new Date();

  const dates = [];
  for (let d = offset; d < offset + 6; d++) {
    const day = new Date();
    day.setDate(today.getDate() - today.getDay() + d + 2);
    dates.push(day);
  }
  return dates;
}
