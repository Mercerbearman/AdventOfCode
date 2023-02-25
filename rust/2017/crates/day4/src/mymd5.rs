use md5;
fn calcDigest() -> String {
    let digest = md5::compute(b"abcde");

    String::new()
}

pub fn myHash() {
    calcDigest();
}