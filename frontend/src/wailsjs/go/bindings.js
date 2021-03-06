// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
const go = {
  "main": {
    "App": {
      /**
       * Connect
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<void>} 
       */
      "Connect": (arg1) => {
        return window.go.main.App.Connect(arg1);
      },
      /**
       * Greet
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<string>}  - Go Type: string
       */
      "Greet": (arg1) => {
        return window.go.main.App.Greet(arg1);
      },
      /**
       * IsAuthorized
       * @returns {Promise<boolean>}  - Go Type: bool
       */
      "IsAuthorized": () => {
        return window.go.main.App.IsAuthorized();
      },
      /**
       * OpenAuthorization
       * @returns {Promise<void>} 
       */
      "OpenAuthorization": () => {
        return window.go.main.App.OpenAuthorization();
      },
      /**
       * SendChatMessage
       * @param {string} arg1 - Go Type: string
       * @param {string} arg2 - Go Type: string
       * @returns {Promise<void>} 
       */
      "SendChatMessage": (arg1, arg2) => {
        return window.go.main.App.SendChatMessage(arg1, arg2);
      },
      /**
       * UserList
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<Array<string>>}  - Go Type: []string
       */
      "UserList": (arg1) => {
        return window.go.main.App.UserList(arg1);
      },
    },
  },

};
export default go;
