use std::fs;
// use std::env;
// use std::path::PathBuf;
use std::io::Error;
use std::collections::HashMap;

const INPUT_FILE : &str = "crates/day1/input.txt";

// //Helper to get current directory without error.
// fn get_cur_dir() -> std::io::Result<PathBuf> {
//     let path = env::current_dir()?;
//     Ok(path)
// }

// //Append 
// fn get_input_path() -> PathBuf {
//     let idx = "day1";
//     let mut pa = get_cur_dir().unwrap();
//     let file_p : String = ["crates/", idx, "/input.txt"].into_iter().collect();
//     pa.push(file_p);
//     pa
// }

//Get data from our input file.
fn get_input() -> Result<String, Error> {
    let line = fs::read_to_string(INPUT_FILE)?;
    Ok(line)
}

//Lets do some magic regarding folding.
// This iterates over input string and 'folds' the characters it finds into a new
// has map.  We declare the has map as mutable as we are modifying it in the fold method directly.
fn process_data_part1(input: &String) -> i32 {
    let char_counts: HashMap<char, i32> = 
        input
            .chars()
            .fold(HashMap::new(), |mut map, c| {
                *map.entry(c).or_insert(0) += 1;
                map
            });
    println!("{:?}", char_counts);

    let floor = char_counts.get(&'(').unwrap_or(&0) - char_counts.get(&')').unwrap_or(&0);

    println!("Floor {}", floor);

    //Return for tests.
    floor
}

//This is a more brute force approach.  We simply iterate over the characters until we hit the basement (negative number).
fn process_data_part2(input: &String) -> i32 {
    let mut floor = 0;
    let mut idx = 1;

    for cha in input.chars().into_iter() {
        match cha {
            '(' => floor += 1,
            ')' => {
                floor -= 1;
                if floor < 0 {
                    println!("We entered basement on position {}", idx);
                    return idx;
                };
            }
            _ => println!("Found illegal char")
        }
        idx += 1;
    }

    idx
}

pub fn day1() -> std::io::Result<()>{
    match get_input() {
        Err(x) => {
            println!("Couldn't read file, got err {}", x);
            //println!("Current directory is {}", get_cur_dir().unwrap().display());
        },
        Ok(inp) => {
            
            println!("Read file: ../input.txt with result...");
            println!("{}", inp);
            process_data_part1(&inp);
            process_data_part2(&inp);
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
    fn test_part1() {
        assert_eq!(0, process_data_part1(&"(())".to_string()));
        assert_eq!(0, process_data_part1(&"()()".to_string()));
        assert_eq!(3, process_data_part1(&"(((".to_string()));
        assert_eq!(3, process_data_part1(&"(()(()(".to_string()));
        assert_eq!(3, process_data_part1(&"))(((((".to_string()));
        assert_eq!(-1, process_data_part1(&"())".to_string()));
        assert_eq!(-1, process_data_part1(&"))(".to_string()));
        assert_eq!(-3, process_data_part1(&")))".to_string()));
        assert_eq!(-3, process_data_part1(&")())())".to_string()));
    }

    #[test]
    fn test_part2() {
        assert_eq!(1, process_data_part2(&")".to_string()));
        assert_eq!(5, process_data_part2(&"()())".to_string()));
    }
}