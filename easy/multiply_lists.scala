object Main extends App {
  val source = scala.io.Source.fromFile(args(0))
  val lines = source.getLines.filter(_.length > 0)

  for (l <- lines) {
    val lists = l.split(""" \| """).toList
    val la = lists(0).split(" ").map(_.toInt)
    val lb = lists(1).split(" ").map(_.toInt)
    var r = new Array[String](la.length)
    for (i <- 0 to (la.length - 1)) {
      r(i) = (la(i) * lb(i)).toString
    }
    println(r.mkString(" "))
  }
}
