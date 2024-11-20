// This is example code
void example() {
  print("How many fibonacci numbers do you want?");
  var nums = 100;
  if (nums <= 0) {
    print("Wrong value");
    return;
  }
  var a = 0;
  var b = 1;

  while (nums > 0) {
    print(a);
    var c = a + b;
    a = b;
    b = c;
    nums = nums - 1;
  }
}
