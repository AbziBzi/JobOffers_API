import React, { createContext, useState } from "react";

export const UserContext = createContext();

// This context provider is passed to any component requiring the context
export const UserProvider = ({ children }) => {
    const [token, setToken] = useState("");
    const [id, setId] = useState(0);
    const [roleId, setRoleId] = useState(0);

    return (
        <UserContext.Provider
            value={{
                token,
                id,
                roleId,
                setToken,
                setId,
                setRoleId
            }}
        >
            {children}
        </UserContext.Provider>
    );
};