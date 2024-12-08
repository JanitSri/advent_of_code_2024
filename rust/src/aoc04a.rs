use crate::util;

const XMAS_ARRAY: [char; 4] = ['X', 'M', 'A', 'S'];

fn aoc04a() -> i32 {
    let contents = util::read_file("../inputs/04.txt");

    let ws: Vec<Vec<char>> = contents.split("\n").map(|f| f.chars().collect()).collect();

    let directions = [
        [-1, 0],  // Up
        [-1, 1],  // Up-Right
        [0, 1],   // Right
        [1, 1],   // Down-Right
        [1, 0],   // Down
        [1, -1],  // Down-Left
        [0, -1],  // Left
        [-1, -1], // Up-Left
    ];

    let mut total = 0;
    for (ri, r) in (0..ws.len() as i32).zip(ws.iter()) {
        for (ci, _) in (0..r.len() as i32).zip(r.iter()) {
            for d in directions {
                if traverse(ri, ci, 0, &ws, d) {
                    total += 1;
                }
            }
        }
    }

    total
}

fn traverse(ri: i32, ci: i32, xi: i32, arr: &Vec<Vec<char>>, d: [i32; 2]) -> bool {
    if ri < 0
        || ri >= arr.len() as i32
        || ci < 0
        || ci >= arr[ri as usize].len() as i32
        || xi >= XMAS_ARRAY.len() as i32
    {
        return false;
    }

    if arr[ri as usize][ci as usize] == XMAS_ARRAY[xi as usize] {
        if xi == XMAS_ARRAY.len() as i32 - 1 {
            return true;
        }

        return traverse(ri + d[0], ci + d[1], xi + 1, arr, d);
    }

    false
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_aoc03b() {
        assert_eq!(2591, aoc04a());
    }
}
