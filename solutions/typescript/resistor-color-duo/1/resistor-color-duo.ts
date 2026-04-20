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

export function decodedValue(colorsArray: Color[]): number {
  const firstColor = colorsArray[0] as Color
  const secondColor = colorsArray[1] as Color
  return Number(COLORS[firstColor] + COLORS[secondColor])
}
