import {FC, createContext, useState, useEffect, ReactNode} from "react";
import axios, {AxiosError} from "axios";
import {ErrorResponse, SuccessResponse} from "../types/response.ts";
import {InitialResponse} from "../types/payload.ts";

export const GlobalContext = createContext<{ state?: InitialResponse, error?: string }>({});

export const GlobalContextProvider: FC<{ children: ReactNode }> = ({children}) => {
    const [state, setState] = useState<InitialResponse | undefined>(undefined);
    const [error, setError] = useState<string | undefined>(undefined);

    const value = {
        state: state,
        error: error,
    };

    useEffect(() => {
        const urlParams = new URLSearchParams(window.location.search);
        const token = urlParams.get("token");
        axios.get<SuccessResponse<InitialResponse>>("/wonderland/api/pair/initial", {
            params: {
                sessionNo: token,
            }
        })
            .then((res) => {
                setState(res.data.data);
            })
            .catch((err: AxiosError<ErrorResponse>) => {
                setError(err.response?.data.message || err.message);
            });
    }, []);

    return (
        <GlobalContext.Provider value={value}>
            {
                children
            }
        </GlobalContext.Provider>
    );
};