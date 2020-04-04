
window.onload = function() {
    const socket = new WebSocket('ws://localhost/api')

    let hasGrid = false

    socket.onopen = function(){
        console.log('Listening')
    }

    socket.onmessage = function(evt){
        var obj = JSON.parse(evt.data)
        handleMessage(obj)
    }

    function handleMessage(message) {
        console.log('received message')
        const game = document.querySelector('#game')

        if (!hasGrid) {
            message.grid.forEach((row, rowIndex) => {
                const rowElement = document.createElement('div')
                rowElement.setAttribute('class', 'row')

                row.cells.forEach((cell, cellIndex) => {
                    let cellElement = null
                    if (cell.kind === 'definition') {
                        cellElement = document.createElement('span')
                        cellElement.textContent = cell.value
                    } else {
                        cellElement = document.createElement('input')
                        cellElement.setAttribute('value', cell.value)
                        cellElement.onchange = sendUpdate
                        cellElement.setAttribute('maxLength', '1')

                    }
                    cellElement.setAttribute('class', 'cell ' + cell.kind)
                    cellElement.setAttribute('id', `id-${rowIndex}-${cellIndex}`)

                    rowElement.append(cellElement)
                })
                game.append(rowElement)
            })

            hasGrid = true
        } else {
            message.grid.forEach((row, rowIndex) => {
                row.cells.forEach((cell, cellIndex) => {
                    if (cell.kind === 'fillable') {
                        document.querySelector(`#id-${rowIndex}-${cellIndex}`).setAttribute('value', cell.value)
                    }
                })
            })
        }

        function sendUpdate(e) {
            console.log('send update !')
            const ids = e.target.id.split('-')
            e.target.value = e.target.value.toUpperCase()
            socket.send(JSON.stringify({
                kind: 'update',
                row: parseInt(ids[1]),
                cell: parseInt(ids[2]),
                value: e.target.value,
            }))
        }
    }
}
