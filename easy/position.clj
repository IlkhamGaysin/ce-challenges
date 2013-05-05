(require '[clojure.string :as str])
(doseq [item (for [st (str/split-lines (slurp (first *command-line-args*)))]
  (#(= (bit-test (Integer/parseInt (first %))
                 (- (Integer/parseInt (nth % 1)) 1))
       (bit-test (Integer/parseInt (first %))
                 (- (Integer/parseInt (nth % 2)) 1)))
(str/split st #",")))] (prn item))
