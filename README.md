# Small dogApp
Goal: Get to know a bit more React, javascript.

## Version

dogApp version 1.0.0
Go 1.20
npm 9.6.0

## Backend

Backend is written in GO. Webserver serves a API at localhost:port where port is the environment variable BACKEND_PORT.
```
export BACKEND_PORT=8080
```
Use of imaginaryDatabase.txt as focus is not on that.
Can be improved to use a postgres in Docker. Read/Write to that database using API requests.

### Run

Instructions:
```
cd backend
GOOS=linux go build .
export BACKEND_PORT=8080
./backend
```

Example:
```
➜  backend ./backend 
2023/03/05 11:43:59 Starting up the webserver at http://localhost:8080/api/dogs
```

## Frontend

Frontend contains HTML, CSS & javascript using webpack.

### Run

Instructions:
```
cd frontend
npm install
npm start
```

Example:
```
➜  frontend git:(master) ✗ npm start

> dogApp@1.0.0 start
> ../node_modules/.bin/webpack serve --mode development --open

<i> [webpack-dev-server] Project is running at:
<i> [webpack-dev-server] Loopback: http://localhost:3000/
<i> [webpack-dev-server] On Your Network (IPv4): http://192.168.0.138:3000/
<i> [webpack-dev-server] On Your Network (IPv6): http://[]:3000/
<i> [webpack-dev-server] Content not from webpack is served from '/path/to/dogApp/frontend/dist' directory
<i> [webpack-dev-middleware] wait until bundle finished: /
asset bundle.js 1.38 MiB [emitted] (name: main)
runtime modules 27.4 KiB 14 modules
modules by path ../node_modules/ 1.25 MiB 36 modules
modules by path ./src/ 5.85 KiB
  modules by path ./src/*.js 2.74 KiB
    ./src/index.js 204 bytes [built] [code generated]
    ./src/App.js 2.54 KiB [built] [code generated]
  modules by path ./src/*.css 3.12 KiB
    ./src/styles.css 2.27 KiB [built] [code generated]
    ../node_modules/css-loader/dist/cjs.js!./src/styles.css 863 bytes [built] [code generated]
webpack 5.75.0 compiled successfully in 1494 ms
```