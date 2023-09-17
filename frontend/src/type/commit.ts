export type Success = {
    success: boolean,
    data: {
        matched : boolean,
        forwardLink: string
      }
};
  
export type Err = {
    success: boolean,
    code: string,
    message: string,
    error: string
  };