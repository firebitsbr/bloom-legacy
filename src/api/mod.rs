mod errors;
mod state;
pub mod middlewares;

use serde::Serialize;
use actix_web::{
    Result as ActixResult,
    HttpResponse,
    HttpRequest,
    dev,
};

pub use errors::Error;
pub use state::State;


#[derive(Serialize)]
pub struct Response<T: Serialize> {
    pub data: Option<T>,
    pub error: Option<ResponseError>
}

#[derive(Serialize)]
pub struct ResponseError{
    message: String,
}

impl<T: Serialize> Response<T> {
    pub fn data(data: T) -> Response<T> {
        return Response{
            data: Some(data),
            error: None,
            };
    }

    pub fn error(err: &str) -> Response<T> {
        return Response{
            data: None,
            error: Some(ResponseError{message: err.to_string()}),
        };
    }
}


pub fn route_404(_req: &HttpRequest<State>) -> ActixResult<HttpResponse> {
    return Err(Error::RouteNotFound.into());
}

pub fn json_default_config(cfg: &mut (dev::JsonConfig<State>, ())) {
    cfg.0.error_handler(|err, _req| {  // <- create custom error response
        Error::BadClientData{error: err.to_string()}.into()
    });
}
