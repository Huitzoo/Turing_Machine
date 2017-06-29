package main

import (
  "fmt"
  "strings"
  "image"
  "image/gif"
  "os"
)
var i int
var files []string
var cadena string

func Estado_A(){
  // 0 -> 0,A> 1-> 0,B,>
  if strings.Compare(string(cadena[i]),"0")==0 {
    files = append(files,"Imagenes/estado_a_avanza.gif")
    files = append(files,"Imagenes/fondo.gif")
    i++
    Estado_A()
  }else{
    fmt.Println("cambiar")
    files = append(files,"Imagenes/estado_a_cambia.gif")
    files = append(files,"Imagenes/fondo.gif")
    cadena = strings.Replace(cadena,string(cadena[i]),"0",1)
    files = append(files,"Imagenes/estado_a_cambiado.gif")
    files = append(files,"Imagenes/cambio.gif")
    fmt.Println(cadena)
    i++
    Estado_B()
  }
}

func Estado_B(){
    // 0 -> 1,@ 1-> 1,B,>
    if strings.Compare(string(cadena[i]),"0")==0{
      var cadena_aux string
      fmt.Println("encontre 0")
      files = append(files,"Imagenes/estado_b_cambia.gif")
      files = append(files,"Imagenes/fondo.gif")
      cadena_aux = cadena[i:len(cadena)]
      cadena_aux1 := cadena[0:i]
      fmt.Println(cadena_aux1)
      cadena_aux = strings.Replace(cadena_aux,string(cadena[i]),"1",1) 
      fmt.Println("soy: ",cadena_aux)
      cadena = ""
      cadena = cadena_aux1 + cadena_aux
      files = append(files,"Imagenes/estado_b_cambiado.gif")
      files = append(files,"Imagenes/final.gif")
      fmt.Println(cadena)
      return
    }else{
      files = append(files,"Imagenes/estado_b_avanza.gif")
      files = append(files,"Imagenes/fondo.gif")
      i++
      Estado_B()
    }
}

func create_gif(titulo string){

  outGif := &gif.GIF{}
    for _, name := range files {
        f, _ := os.Open(name)
        inGif, _ := gif.Decode(f)
        f.Close()
        outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
        outGif.Delay = append(outGif.Delay, 90)
    }

    // save to out.gif
    f, _ := os.OpenFile(titulo, os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    gif.EncodeAll(f, outGif)
}

func validar(){
  var bandera,j,k int = 0,0,0
  fmt.Println("cadnea: ",cadena)
  for n := 0; n < len(cadena); n++ {
    if strings.Compare(string(cadena[n]),"0") == 0{
      files = append(files,"Imagenes/0.gif")
    }else{
      files = append(files,"Imagenes/1.gif")
    }
    files = append(files,"Imagenes/guion.gif")
  }
  for j = 0; j < len(cadena); j++ {
    if strings.Compare(string(cadena[j]),"1") == 0 {  
      break 
    }
  }
  for k = j ; k<len(cadena);k++{
    if strings.Compare(string(cadena[k]),"0")==0{
      if k == len(cadena)-1 || k == len(cadena)-2{
          bandera = 0
        }else{
          files = append(files,"Imagenes/error.gif")
          bandera = 1
          break
        }
      }else{
        bandera = 0
    }
  }
  if bandera == 0{
    print("Resultado: ", k-j-1)
    files = append(files,"Imagenes/valida.gif")
  }
}

func main()  {
  i=0
  fmt.Println("Suma de dos numeros: ")
  fmt.Println("Ingresa una cadena: 01011")
  fmt.Scanln(&cadena)
  for j := 0; j < len(cadena); j++ {
    print(string(cadena[j]))
    if strings.Compare(string(cadena[j]),"1") != 0 {  
      if strings.Compare(string(cadena[j]),"0") != 0 {
      } 
    }
  }
  Estado_A()
  //trancisiones
  create_gif("trancisiones.gif")
  files = nil
  validar()
  //cadena final valida o no
  create_gif("resultado.gif")
}
