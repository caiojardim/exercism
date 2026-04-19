const COLORS = {
  "black": "0",
  "brown": "1",
  "red": "2",
  "orange": "3",
  "yellow": "4",
  "green": "5",
  "blue": "6",
  "violet": "7",
  "grey": "8",
  "white": "9",
} as const

type Color = keyof typeof COLORS

function formatNumberToText (number: number): string {
  if (number >= 1_000_000_000) {
    return `${number / 1_000_000_000} gigaohms`
  }
  if (number >= 1_000_000) {
    return `${number / 1_000_000} megaohms`
  }
  if (number >= 1_000) {
    return `${number / 1_000} kiloohms`
  }
  return `${number} ohms`
}

export function decodedResistorValue([first, second, third]: Color[]): string {
  const parsedNumber = Number(COLORS[first] + COLORS[second] + "0".repeat(Number(COLORS[third])))
  return formatNumberToText(parsedNumber)
}
