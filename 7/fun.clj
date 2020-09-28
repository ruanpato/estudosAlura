(defn imprime-mensagem []
    (println "----------------------------")
    (println "Uma mensagem")
)

;;(defn aplica-desconto [valor-bruto, desconto]
;;    (* valor-bruto desconto)
;;)

;;(defn valor-descontado [valor-bruto]
;;    (println "Desconto de" (- 1.0 0.9) "%")
;;    (aplica-desconto valor-bruto 0.9)  
;;)

(defn valor-descontado
    "Retorna o valor com desconto de 10% se o valor bruto for estritamente maior que 100."
    [valor-bruto]
    (if (> valor-bruto 100)
      (let [taxa-de-desconto (/ 10 100)
            desconto (* valor-bruto taxa-de-desconto)]
        (println "Calculando desconto de" desconto)
        (- valor-bruto desconto))
        valor-bruto))

(valor-descontado 101)

(- 1 0.90)

(class 90N)