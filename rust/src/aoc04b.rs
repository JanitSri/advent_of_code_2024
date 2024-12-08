use crate::util;

const XMAS_ARRAY: [char; 4] = ['X', 'M', 'A', 'S'];

fn aoc04b() -> i32 {
    let contents = util::read_file("../inputs/04.txt");

    let ws: Vec<Vec<char>> = contents.split("\n").map(|f| f.chars().collect()).collect();

    let mut total = 0;
    for (ri, r) in (0..ws.len() as i32).zip(ws.iter()) {
        for (ci, _) in (0..r.len() as i32).zip(r.iter()) {
            if ws[ri as usize][ci as usize] == 'A' {
                let d1 = (traverse(ri - 1, ci - 1, 'M', &ws) && traverse(ri + 1, ci + 1, 'S', &ws))
                    || traverse(ri - 1, ci - 1, 'S', &ws) && traverse(ri + 1, ci + 1, 'M', &ws);
                let d2 = (traverse(ri - 1, ci + 1, 'M', &ws) && traverse(ri + 1, ci - 1, 'S', &ws))
                    || (traverse(ri - 1, ci + 1, 'S', &ws) && traverse(ri + 1, ci - 1, 'M', &ws));

                if d1 && d2 {
                    total += 1;
                }
            }
        }
    }

    total
}

fn traverse(ri: i32, ci: i32, c: char, arr: &Vec<Vec<char>>) -> bool {
    if ri < 0 || ri >= arr.len() as i32 || ci < 0 || ci >= arr[ri as usize].len() as i32 {
        return false;
    }

    arr[ri as usize][ci as usize] == c
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_aoc03b() {
        assert_eq!(1880, aoc04b());
    }
}
