use core::str;
use std::fs;

pub fn read_file(path: &str) -> String {
    let contents = fs::read(path).expect("cannot read file {path}");

    String::from_utf8(contents).expect("could not convert to string")
}