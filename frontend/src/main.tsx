import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import {ChakraProvider} from "@chakra-ui/react";
import {GlobalContextProvider} from "./contexts/GlobalState.tsx";

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <ChakraProvider>
            <GlobalContextProvider>
                <App/>
            </GlobalContextProvider>
        </ChakraProvider>
    </React.StrictMode>,
)
