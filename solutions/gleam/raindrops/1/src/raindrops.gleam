import gleam/int
pub fn convert(number: Int) -> String {
  case add(number, 3, "Pling") 
    <> add(number, 5, "Plang")
    <> add(number, 7, "Plong") {
    "" -> int.to_string(number)
    i -> i
  } 
}

fn add(number, div, sound) {
  case number%div {
    0 -> sound
    _ -> ""
  }
}