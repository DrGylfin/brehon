
class Reactive {
	
	constructor(value) {
		this._val = value;
		this._subscribers = [];
		
	}
	
	get value() {
		return this._val;
	}
	
	set value(newValue) {
		this._val = newValue;
		for (const subscribeFunction of this._subscribers) {
			subscribeFunction(newValue);
		}
	}
	
	subscribe(func) {
		this._subscribers.push(func);
		func(this._val);
	}
	
}

let todoItems = [];
let items = new Reactive(todoItems);
items.subscribe(() => {
	renderTodoList(items.value);
});

function chooseWisely() {
	
	let itemCount = items.value.length;
	if (itemCount == 0) {
		return undefined;
	} else if (itemCount == 1) {
		return -1;
	} else {
		
		// choose a random number between 0 and one less than itemCount
		// since the list of items is 0-indexed
		return Math.floor(Math.random() * itemCount);
		
	}
	
}


/********
TODO: allow editing
********/
function showItemInfo(item) {
	
	let itemText = items.value[item].text;
	let itemCat = items.value[item].category;
	let itemTime = items.value[item].timereq;
	
	document.getElementById('item-text').value = itemText;
	document.getElementById('item-cat').value = itemCat;
	document.getElementById('item-time').value = itemTime;

}

function getItemsFromFile() {

	let promisedData = window.go.main.FileHandler.LoadList().then((data, err) => {
		items.value = JSON.parse(data);
	});

}

/********
Javascript function for sending the todo list items to the Go backend for saving.
*********/
function saveTodoList() {
	let result = window.go.main.FileHandler.SaveList(items.value).then((str, err) => {
		if (err != undefined) {
			console.log(err);
		}
	});
}


function addTodoItem() {
	
	let newText = document.getElementById('item-text');
	let newCat = document.getElementById('item-cat');
	let newTime = document.getElementById('item-time');

	if (newText.value == '' || newCat.value == '' || newTime.value == '') {
		
		console.log('Missing input data');
		
	} else {
		
		let newId = items.value.length;
		items.value = [...items.value, {"id":newId, "text":newText.value, "category":newCat.value, "timereq":newTime.value}];
		newText.value = '';
		newCat.value = '';
		newTime.value = '';
		
	}
	
}

function deleteTodoItem(item) {
	
	let index = item.id;
	let currentData = items.value;
	let removed = currentData.splice(index, 1);				// TODO: undo function
	items.value = currentData;
	
	// clear input boxes, just in case
	document.getElementById('item-text').value = '';
	document.getElementById('item-cat').value = '';
	document.getElementById('item-time').value = '';
}

function renderTodoList(data = items.value) {
	
	let listHTML = '';
	data.forEach((item, index, array) => {
		listHTML += '<li id="' + index + '">' + data[index].text + iconDelete + '</li>';
	});
	todoList.innerHTML = listHTML;
	
}

window.addEventListener('unload', saveTodoList, false);