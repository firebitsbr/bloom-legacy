use crate::{
    api,
    accounts::api::v1::models,
    log::macros::*,
    accounts::controllers,
    api::middlewares::{
        GetRequestLogger,
        GetRequestId,
        GetRequestAuth,
    },
    error::KernelError,
};
use futures::future::Future;
use actix_web::{
    web, Error, HttpRequest, HttpResponse,
};
use futures::future::IntoFuture;
use rand::Rng;
use std::time::Duration;


pub fn post(data: web::Json<models::PasswordResetRequestBody>, state: web::Data<api::State>, req: HttpRequest)
-> impl Future<Item = HttpResponse, Error = Error> {
    let logger = req.logger();
    let request_id = req.request_id().0;
    let mut rng = rand::thread_rng();
    let config = state.config.clone();
    let auth = req.request_auth();

    let session_id = match auth.session {
        Some(ref session) => Some(session.id),
        None => None,
    };

    // random sleep to prevent bruteforce and sidechannels attacks
    return tokio_timer::sleep(Duration::from_millis(rng.gen_range(250, 350))).into_future()
    .from_err()
    .and_then(move |_|
        state.db
        .send(controllers::RequestPasswordReset{
            config,
            email_or_username: data.email_or_username.clone(),
            request_id,
            session_id,
        }).flatten()
    )
    .and_then(move |_| {
        let res = api::Response::data(models::NoData{});
        Ok(HttpResponse::Ok().json(&res))
    })
    .map_err(move |err: KernelError| {
        slog_error!(logger, "{}", err);
        return err;
    })
    .from_err()
    .responder();
}
