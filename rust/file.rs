use std::io::Write;

fn main() {
    let mut f = std::fs::File::create("foo.txt").expect("Couldn't create foo.tx");
    let buf = b"Hello file";
    let result = f.write(buf);
}
