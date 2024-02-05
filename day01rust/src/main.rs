use std::{fs::File, io::{self, BufRead}, path::Path};

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {
    let filename = "./input";

    let mut inputs = Vec::new();
    
    if let Ok(lines) = read_lines(filename) {
        for line in lines.flatten() {
            inputs.push(line)
        }
    }

    // Part 1
    let mut result1 = 0;
    for input in &inputs {
        
        let first_number = input.chars().find(|x| x.is_digit(10));
        let x1 = first_number.unwrap_or_else(|| ' ');
        let last_number = input.chars().rev().find(|x| x.is_digit(10));
        let x2 = last_number.unwrap_or_else(|| ' ');

        result1 = result1 + create_number(x1, x2);
    }

    println!("Part1: {}", result1);


    // Part 2
    let mut result2 = 0;
    let digits_in_words = vec!{
        ("one", '1'),
        ("two", '2'),
        ("three", '3'),
        ("four", '4'),
        ("five",'5'),
        ("six",'6'),
        ("seven",'7'),
        ("eight",'8'),
        ("nine",'9'),
    };

    for input in &inputs {

        let vec_chars: Vec<char> = input.chars().collect();
        // First numbers
        let first = vec_chars.iter().position(|x| x.is_digit(10));
        let (mut first_position, mut first_value)= match first {
            Some(n) => { (n, vec_chars[n])  },
            None => (vec_chars.len()+1, ' '),
        };
        // Last numbers

        let last = vec_chars.iter().rposition(|x: &char| x.is_digit(10));
        let (mut last_position, mut last_value) = match last {
            Some(n) => (n, vec_chars[n]),
            None => (0, ' '),
        };

        for digit_in_word in digits_in_words.iter() {
            let first_digit = input.find(digit_in_word.0);
            (first_position, first_value) = match first_digit {
                Some(number) => {
                    let (result, value) = if first_position > number {
                        (number, digit_in_word.1)
                    } else {
                        (first_position, first_value)
                    };
                    (result, value)

                },
                None => (first_position, first_value),
            };

            // Last

            let last_digit = input.rfind(digit_in_word.0);
            (last_position, last_value) = match last_digit {
                Some(number) => {
                    let (result, value) = if number >= last_position {
                        (number, digit_in_word.1)
                    } else {
                        (last_position, last_value)
                    };
                    (result, value)

                },
                None => (last_position, last_value),
            };

            // let last_number = input.chars().rev().find(|x| x.is_digit(10)); 
            
            
        }

        result2 = result2 + create_number(first_value, last_value);
        
    }
    println!("Part2: {}", result2);


}

fn create_number(x1: char, x2: char) -> i32 {
    let string_number = format!("{}{}", x1, x2);
    string_number.parse::<i32>().unwrap_or_else(|_| 0)
}
