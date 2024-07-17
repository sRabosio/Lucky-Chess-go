/**@type {(el:HTMLElement, styles:object.<string,string>)=>void} */
const applyStyles = (el, styles)=>{
  Object.entries(styles)
  .forEach(entry=>{
    el.style[entry.at(0)] = entry.at(1)
  })
}


/**@type {(ms: number)=>Promise<void>} */
const wait = (ms)=>new Promise((resolve)=>{
  setTimeout(resolve, ms)
})