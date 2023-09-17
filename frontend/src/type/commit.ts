type Success = {
    success: boolean,
    data: {
        matched : boolean,
        forwardLink: string
      }
};
  
  type Err = {
    success: boolean,
    code: String,
    message: String,
    error: String
  };