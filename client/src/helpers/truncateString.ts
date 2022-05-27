export default function truncateString(str: string | null, num: number) {
  if (str && str.length > num) {
    return str.slice(0, num) + "...";
  } else {
    return str;
  }
}
