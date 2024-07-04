use std::str::FromStr;
pub fn is_armstrong_number(num: u32) -> bool {
    let big_num: u64 = u64::from(num);
    let num_string = String::from_str(&big_num.to_string()).expect("success");
    let num_char_length: u32 = num_string.len().try_into().unwrap();
    let num_array: Vec<u64> = num_string
        .chars()
        .map(|s| u64::from(s.to_digit(10).expect("")))
        .map(|s: u64| s.pow(num_char_length))
        .collect();
    println!("{:?}", num_array);
    let armstrong_num: u64 = num_array.into_iter().sum();
    //     .expect("Error reducing");
    println!("{}", armstrong_num);
    u64::from(num) == armstrong_num
}
