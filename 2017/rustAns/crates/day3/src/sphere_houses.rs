use std::fs;
use std::io::Error;
use std::collections::HashMap;

const INPUT_FILE : &str = "crates/day3/input.txt";

//Get data from our input file.
fn get_input() -> Result<String, Error> {
    let line = fs::read_to_string(INPUT_FILE)?;
    Ok(line)
}

#[derive(Eq, Hash, PartialEq, Clone, Copy)]
struct Point {
    x: i32,
    y: i32
}

impl Point {
    fn equals(self: &Self, point: &Point) -> bool {
        self.x == point.x && self.y == point.y
    }
}

//Lets do some magic regarding folding.
// This iterates over input string and 'folds' the characters it finds into a new
// has map.  We declare the has map as mutable as we are modifying it in the fold method directly.
fn part1(input: &String) -> i32 {
    let mut locs : HashMap<Point, i32> = HashMap::new();

    let mut currPoint = Point{x: 0, y: 0};

    *locs.entry(currPoint).or_insert(0) += 1;

    for sym in input.chars() {
        match sym {
            '^' => currPoint.y += 1,
            'v' => currPoint.y -= 1,
            '>' => currPoint.x += 1,
            '<' => currPoint.x -= 1,
            _ => println!("Bad input symbol!")
        };

        *locs.entry(currPoint).or_insert(0) += 1;
    }

    let v : Vec<_> = locs.values().filter(|&x| x >= &1).collect();//.filter(|h| h >= 1).collect();

    return v.len() as i32;
}

//This is a more brute force approach.  We simply iterate over the characters until we hit the basement (negative number).
fn part2(input: &String) -> i32 {
    let mut locs : HashMap<Point, i32> = HashMap::new();

    let mut currPointSanta = Point{x: 0, y: 0};
    let mut currPointRobo = Point{x: 0, y: 0};

    *locs.entry(currPointSanta).or_insert(0) += 1;
    *locs.entry(currPointRobo).or_insert(0) += 1;

    let santasIter :Vec<_> = input.chars().enumerate()
        .map(|(i,sym)|{
            if i % 2 == 0 {
                //Santa portion..
                match sym {
                    '^' => currPointSanta.y += 1,
                    'v' => currPointSanta.y -= 1,
                    '>' => currPointSanta.x += 1,
                    '<' => currPointSanta.x -= 1,
                    _ => println!("Bad input symbol!")
                };

                *locs.entry(currPointSanta).or_insert(0) += 1;
            }
            else {
                //Robot portion..
                match sym {
                    '^' => currPointRobo.y += 1,
                    'v' => currPointRobo.y -= 1,
                    '>' => currPointRobo.x += 1,
                    '<' => currPointRobo.x -= 1,
                    _ => println!("Bad input symbol!")
                };

                *locs.entry(currPointRobo).or_insert(0) += 1;
            }
        }).collect();

    let v : Vec<_> = locs.values().filter(|&x| x >= &1).collect();//.filter(|h| h >= 1).collect();

    return v.len() as i32;
}

pub fn houses() -> std::io::Result<()>{
    match get_input() {
        Err(x) => {
            println!("Couldn't read file, got err {}", x);
            //println!("Current directory is {}", get_cur_dir().unwrap().display());
        },
        Ok(inp) => {
            
            println!("Read file: ../input.txt with result...");
            println!("{}", inp);
            let ret1 = part1(&inp);
            println!("Part1: {}", ret1);
            let ret2 = part2(&inp);
            println!("Part2: {}", ret2);
        }
    }
    Ok(())
}

//Tests.
#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_part1_1() {
        assert_eq!(2, part1(&">".to_string()));
    }

    #[test]
    fn test_part1_2() {
        assert_eq!(4, part1(&"^>v<".to_string()));
    }

    #[test]
    fn test_part1_3() {
        assert_eq!(2, part1(&"^v^v^v^v^v".to_string()));
    }

    #[test]
    fn test_part2_1() {
        assert_eq!(3, part2(&"^v".to_string()));
    }

    #[test]
    fn test_part2_2() {
        assert_eq!(3, part2(&"^>v<".to_string()));
    }

    #[test]
    fn test_part2_3() {
        assert_eq!(11, part2(&"^v^v^v^v^v".to_string()));
    }
}