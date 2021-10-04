# EjercicioGo Trabajo Practico TUDAI 2021 (Test coverage: 87.1% of statements)


Para la resolucion del ejercicio implemente dos funciones.
`TransformarCadena` la cueal recibe por parametro el string a evaluar y devuelve una estructura del tirpo Result.
Esta funcion se encarga de generar la estructura con los valores iniciales, separando segun lo esperado en sus tres tipos de atributos.
type Result struct {
    Type    string 
    Value   string
    Length  int
}
requisitos:
- los primeros dos caracteres son el tipo
- los segundos dos caracteres son el largo del valor
- los siguientes N caracteres son el valor, donde N es el valor del bullet anterior.

Una vez realizado ese paso llamamos a la funcion `validarObjeto` que recibe esta estructura antes mencionada y retorna un booleano.
Esta funcion se encarga de validar si la estructura generada por la primer funcion descripta, cumple con los requisitos mencionados arriba y ademas cumpla con que si la cadena comienza con TX el Value correspondera a valores de tipo String, si la cadena comienza con NN el Value correspondera a valores de tipo entero.
En caso de cumplir con los requsitos retornara [true] lo que hace que la funcion principal `TransformarCadena` retorne la estructura como la genero inicialmente. Caso contrario si `validarObjeto` retornara [false] la funcion principal `TransformarCadena` cambia los valores de la estructura a vacio, vacio, 0("","",0) y retorna la misma.