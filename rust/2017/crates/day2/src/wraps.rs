use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

const INPUT_FILE: &str = "crates/day2/input.txt";

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

//Get data from our input file.
fn get_input() -> Vec<String> {
    let mut v: Vec<String> = vec![];
    if let Ok(lines) = read_lines(INPUT_FILE) {
        // Consumes the iterator, returns an (Optional) String
        for line in lines {
            if let Ok(present) = line {
                v.push(present)
            }
        }
    }
    v
}

#[derive(Debug, Clone, Copy)]
struct Present {
    length : i32,
    width: i32,
    height: i32
}

impl Present {
    fn find_smallest_side_area(self : &Self) -> i32 {
        let areas = vec![&self.length * &self.width, &self.width * &self.height, &self.height * &self.length];

        *areas.iter().min().unwrap()
    }

    //What a redicilous way to calculate perimeters...
    fn find_smallest_perimeter(self: &Self) -> i32 {
        let lengths = vec![self.length, self.width, self.height];

        let perm = lengths.clone().into_iter().fold(0, |acc, side| acc + 2*side) - *lengths.iter().max().unwrap() * 2;

        perm
    }

    fn area(self: &Self) -> i32 {
        self.height * self.length * self.width
    }
}

fn get_present_area(present: &Present) -> i32 {
    2 * present.length * present.height +
    2 * present.height * present.width + 
    2 * present.width * present.length +
    present.find_smallest_side_area()
}

fn get_ribbon_length(present: &Present) -> i32 {

    present.find_smallest_perimeter() + present.area()
}

fn part1(presents: &Vec<Present>) -> i32 {
    let area = presents.into_iter().fold(0, |acc, p| acc + get_present_area(&p));
    println!("Total Area is {}", area);
    area
}

fn part2(presents: &Vec<Present>) -> i32 {
    let ribbon_length = presents.into_iter().fold(0, |acc, p| acc + get_ribbon_length(&p));
    println!("Total Ribbon Length is {}", ribbon_length);
    ribbon_length
}

pub fn start() {

    let lines = get_input();
    let mut presents: Vec<Present> = vec![];

    // println!("{:?}", lines);

    for line in lines {

        let split: Vec<_> = line.split("x").collect();
        // println!("Line {}", line);
        let present = Present{length: split[0].parse().unwrap(), width: split[1].parse().unwrap(), height: split[2].parse().unwrap()};

        presents.push(present);

  //      println!("Presents: {:?}", presents);

//        let small_side = present.clone().find_smallest_side_area();

        // println!("Small side == {}", small_side);

        // break;

    }

    part1(&presents);
    part2(&presents);
}

//Tests.
#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_part1_1() {
        let v: Vec<Present> = vec![Present{length: 2, width: 3, height: 4}];
        assert_eq!(58, part1(&v));
    }

    #[test]
    fn test_part1_2() {
        let v: Vec<Present> = vec![Present{length: 1, width: 1, height: 10}];
        assert_eq!(43, part1(&v));
    }

    #[test]
    fn test_part2_1() {
        let v: Vec<Present> = vec![Present{length: 2, width: 3, height: 4}];
        assert_eq!(34, part2(&v));
    }

    #[test]
    fn test_part2_2() {
        let v: Vec<Present> = vec![Present{length: 1, width: 1, height: 10}];
        assert_eq!(14, part2(&v));
    }
}
