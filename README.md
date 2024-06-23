# neko
This app uses WebRTC to stream a desktop inside of a docker container. Client can be found here: [demodesk/neko-client](https://github.com/demodesk/neko-client).

For **community edition** neko with GUI and _plug & play_ deployment visit [m1k1o/neko](https://github.com/m1k1o/neko).

### **m1k1o/neko** vs **demodesk/neko**, why do we have two of them?

This project started as a fork of [m1k1o/neko](https://github.com/m1k1o/neko). But over time, development went way ahead of the original one in terms of features, updates and refactoring. The goal is to rebase [m1k1o/neko](https://github.com/m1k1o/neko) repository onto this one and move all extra features (such as chat and emotes) to a standalone plugin.

- This project is aimed to be the engine providing foundation for all applications that are streaming desktop environment using WebRTC to the browser.
- [m1k1o/neko](https://github.com/m1k1o/neko) is meant to be self-hosted replacement for [rabb.it](https://en.wikipedia.org/wiki/Rabb.it): Community edition with well-known GUI, all the social functions (such as chat and emotes) and easy deployment.

Notable differences to the [m1k1o/neko](https://github.com/m1k1o/neko) are:

- Go plugin support.
- Multiple encoding qualities simulcast.
   - Bandwidth estimation and adaptive quality.
- Custom screen size (with automatic sync).
- Single cursor for host - cursor image proxying.
- Custom cursor style/badge for participants.
- Inactive cursors (participants that are not hosting).
- Fallback mode and reconnection improvements:
  - Watching using screencasting.
  - Controlling using websockets.
- Members handling:
  - Access control (view, interactivity, clipboard).
  - Posibility to add external members providers.
  - Persistent login (using cookies).
- Drag and drop passthrough.
- File upload passthrough (experimental).
- Microphone passthrough.
- Webcam passthrough (experimental).
- Bi-directional text/html clipboard.
- Keyboard layouts/variants.
- Metrics and REST API.

## Docs

*TBD.*

# neko-client
Connect to [demodesk/neko](https://github.com/demodesk/neko) backend with self contained vue component. 

For **community edition** neko with GUI and _plug & play_ deployment visit [m1k1o/neko](https://github.com/m1k1o/neko).

## Installation
Code is published to public NPM registry and GitHub npm repository.

```bash
# npm command
npm i @demodesk/neko
# yarn command
yarn add @demodesk/neko
```

### Build

You can set keyboard provider at build time, either `novnc` or the default `guacamole`.

```bash
# by default uses guacamole keyboard
npm run build
# uses novnc keyboard
KEYBOARD=novnc npm run build
```

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Run End-to-End Tests with [Playwright](https://playwright.dev)

```sh
# Install browsers for the first run
npx playwright install

# When testing on CI, must build the project first
npm run build

# Runs the end-to-end tests
npm run test:e2e
# Runs the tests only on Chromium
npm run test:e2e -- --project=chromium
# Runs the tests of a specific file
npm run test:e2e -- tests/example.spec.ts
# Runs the tests in debug mode
npm run test:e2e -- --debug
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```


### Example
API consists of accessing Vue reactive state, calling various methods and subscribing to events. Simple usage:

```html
<!-- import vue -->
<script src="https://unpkg.com/vue"></script>

<!-- import neko -->
<script src="./neko.umd.js"></script>
<link rel="stylesheet" href="./neko.css">

<div id="app">
  <neko ref="neko" server="http://127.0.0.1:3000/api" autologin autoplay />
</div>

<script>
new Vue({
  components: { neko },
  mounted() {
    // access state
    // this.$refs.neko.state.session_id
  
    // call methods
    // this.$refs.neko.setUrl('http://127.0.0.1:3000/api')
    // this.$refs.neko.login('username', 'password')
    // this.$refs.neko.logout()
  
    // subscribe to events
    // this.$refs.neko.events.on('room.control.host', (id) => { })
  },
}).$mount('#app')
</script>
```
