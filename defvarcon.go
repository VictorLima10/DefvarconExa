package main//main es el principal de todo programa y debe tener declarada la función main donde se inicia la ejecución.
import "fmt"//paquete para imprimir
func main(){

/*
//************** variables *********los bool son ellse o true ,cadena string enteros var
var cadena a, b, c edad_int
var cadena  ReadString
var bandera bool
var cadenas [] string para mi teclado(alt [91])
:= con este no necesario declarar tipo de variable
*/
Nombre, Ap_Pater, Ap_Mater, Edad, Cuota, Acreditado := "victor hugo", "Lima", "Ramirez", 15 , 1500.00, true
//**************arreglos **************=declara var, nombre ,tamaño [estatico], espesificar que tipo de dato metemos, espesificar que datos tenemos {}
var mes =[12] string { "Ennero", "Febrero", "Marzo", "Abril", "Mayo", "Julio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre" }
AumentoCuota:=0.0
//*****************condicions**********
    if Acreditado == true{
      if Cuota > 1499{
        fmt.Println("\n  Nombre alumno:",Nombre, Ap_Pater, Ap_Mater,",", "Edad",Edad,"\n  Tu cuota es de:",Cuota,"Acreditacion =", Acreditado,"\n")
      }else{
          fmt.Println("La cuota es incorrecta")
      }
    }else{
        fmt.Println("No acreditas tu cuota es incorrecta")
    }

// *******************ciclos go **************

//cuantos caben en el arreglo(len(mes...))

    for i:=1;      i<=len(mes);     i++ {
        //  for i:=1; i<=13; i++ {
      if i==1{
                AumentoCuota=Cuota
      }else{
           AumentoCuota=(AumentoCuota*0.05)+AumentoCuota
      }
//% f punto decimal pero sin exponente, p. ej. 123.456
//cap nos devuelve la capacidad
//len ver la longitud

/*
//variablenew:=len(mes)
//fmt.Println([inicio:final(excluyen)])
//fmt.Println(nombres[0:final(excluyen)])
//mes[0:2]
*/
        case 1:
          //diferentes formas de mandar a imprimir
        // fmt.Println(mes[i-1]+fmt.Sprintf("        %5.3f",Cuota))
          //  fmt.Println(mes[0] ,"cuota" ,fmt.Sprintf("     %5.3f",AumentoCuota))
          // fmt.Println(mes[0:1] ,fmt.Sprintf("     %5.3f",AumentoCuota))
          // fmt.Println("Enero "+fmt.Sprintf("         %5.3f",AumentoCuota))
        case 2:
        fmt.Println(mes[i-1]+fmt.Sprintf("            %5.3f",AumentoCuota))
        case 3:
          fmt.Println(mes[i-1]+fmt.Sprintf("            %5.3f",AumentoCuota))
        case 4:
            fmt.Println(mes[i-1]+fmt.Sprintf("            %5.3f",AumentoCuota))
        case 5:
            fmt.Println(mes[i-1]+fmt.Sprintf("          %5.3f",AumentoCuota))
        case 6:
            fmt.Println(mes[i-1]+fmt.Sprintf("              %5.3f",AumentoCuota))
        case 7:
           fmt.Println(mes[i-1]+fmt.Sprintf("          %5.3f",AumentoCuota))
        case 8:
            fmt.Println(mes[i-1]+fmt.Sprintf("         %5.3f",AumentoCuota))
        case 9:
            fmt.Println(mes[i-1]+fmt.Sprintf("      %5.3f",AumentoCuota))
        case 10:
           fmt.Println(mes[i-1]+fmt.Sprintf("         %5.3f",AumentoCuota))
        case 11:
            fmt.Println(mes[i-1]+fmt.Sprintf("       %5.3f",AumentoCuota))
        case 12:
            fmt.Println(mes[i-1]+fmt.Sprintf("       %5.3f",AumentoCuota))
      }



    }

/*nil valor nulo variable como valor nulo por que go siempre poe 0

************************* Slices ****************** para redimensionar y no definir el tamaño y tiempo de ejecucion
puntero al arreglo----longitud del arreglo al que apunta -- capacidad


slice:= mes [2:15]
//fmt.Println("\n slice para dividir arreglo",slice,"\n")


//make y append=agregar nuevo elemento
//cap nos devuelve la capacidad
//len ver la longitud
fmt.Println(cap(slice))
copia := make ([]string, len(slice))
copy(copia,slice)
fmt.Println("\n",slice)
fmt.Println(copia)
*/
}
