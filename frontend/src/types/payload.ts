export type InitialResponse = {
    sessionId: string,
    email: string,
    pictures: string[],
}

export type CommitResponse = {
    matched: boolean,
    forwardLink: string,
};
