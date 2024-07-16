use std::fmt::{self, Debug, Display};

pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Clock {
        let formatted_hours: i32 = if hours >= 24 { hours % 24 } else { hours };
        let formatted_minutes: i32 = if minutes >= 60 { minutes % 60 } else { minutes };
        Clock {
            hours: formatted_hours,
            minutes: formatted_minutes,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        todo!("Add {minutes} minutes to existing Clock time");
    }
}

impl Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}

impl Debug for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "")
    }
}

impl PartialEq for Clock {
    fn eq(&self, other: &Self) -> bool {
        (self.hours == other.hours) && (self.minutes == other.minutes)
    }
}
