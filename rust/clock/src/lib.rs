use std::fmt::{self, Display};

pub struct Clock;

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        todo!("Construct a new Clock from {hours} hours and {minutes} minutes");
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        todo!("Add {minutes} minutes to existing Clock time");
    }
}

impl Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "")
    }
}
