export interface go {
  "main": {
    "FileHandler": {
		LoadList():Promise<string|Error>
		SaveList(arg1:RawMessage):Promise<string|Error>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
