window.addEventListener('load', ()=>{
    createButton()
})

const wasm = () => {
    const go = new Go();
    console.log("Cargando el archivo main.wasm");
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
        .then((result) => {
            console.log("Carga completa");
            go.run(result.instance);

        })
}
const createButton = () => {
    let btn = document.createElement('button'),
        container = document.getElementById("container")
    container.innerHTML = ""
    container.appendChild(btn)
    btn.innerHTML = "Iniciar Juego"
    btn.addEventListener('click', ()=>{
        wasm()
    })

}