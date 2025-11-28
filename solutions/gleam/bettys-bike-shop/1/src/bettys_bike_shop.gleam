// TODO: import the `gleam/int` module
// TODO: import the `gleam/float` module
// TODO: import the `gleam/string` module
import gleam/float.{to_string}
import gleam/int.{to_float}
pub fn pence_to_pounds(pence: Int) {
  to_float(pence) /. 100.0
}

pub fn pounds_to_string(pounds) {
  "£" <> to_string(pounds)
}
