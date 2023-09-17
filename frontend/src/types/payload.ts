import {SuccessResponse} from "./response.ts";

export type InitialResponse = SuccessResponse<{
    sessionId: string,
    email: string,
    matched: boolean,
    pictures: string[],
}>

export type CommitResponse = SuccessResponse<{
    matched: boolean,
    forwardLink: string,
}>;
