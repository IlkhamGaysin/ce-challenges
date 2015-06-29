object Main extends App {
  val source = scala.io.Source.fromFile(args(0))
  val lines = source.getLines.filter(_.length > 0)

  for (l <- lines) {
    val m = l.split(" ").foldLeft("")((x: String, y: String) => if (y.length > x.length) y else x)
    var r = new Array[String](m.length)
    for (i <- 0 to (m.length - 1)) {
      r(i) = "*" * i + m(i)
    }
    println(r.mkString(" "))
  }
}
