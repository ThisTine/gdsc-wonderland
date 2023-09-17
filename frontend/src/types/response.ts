export type SuccessResponse<T> = {
    success: boolean,
    data: T
};

export type ErrorResponse = {
    success: boolean,
    code: string,
    message: string,
    error: string
};