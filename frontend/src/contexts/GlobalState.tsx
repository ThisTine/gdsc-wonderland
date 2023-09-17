import {FC, createContext, useState} from "react";

export const StateContext = createContext({});

const getState = () => {
    const token = Cookies.get("token");
    if (token) {
        const decoded = jwtDecode(token);
        return {
            avatar: decoded.avatar,
            email: decoded.email,
            name: decoded.name,
            permission:
                decoded.permission === "admin"
                    ? 2
                    : decoded.permission === "staff"
                        ? 1
                        : 0,
        };
    } else {
        return {
            avatar: "",
            email: "",
            name: "",
            permission: 0,
        };
    }
};

export const StateContextProvider: FC<any> = ({children}) => {
    const [state, setState] = useState(null);

    const value: any = {
        state: state,
    };

    return (
        <StateContext.Provider value={value}>
            {children}
        </StateContext.Provider>
    );
};