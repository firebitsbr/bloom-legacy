use crate::{
    api,
    log::macros::*,
    accounts::controllers,
    accounts::api::v1::models,
    api::middlewares::{
        GetRequestLogger,
        GetRequestAuth,
    },
    error::KernelError,
};
use futures::{
    future,
    future::Future,
};
use actix_web::{
    web, Error, HttpRequest, HttpResponse,
};


pub fn get(state: web::Data<api::State>, req: HttpRequest) -> impl Future<Item = HttpResponse, Error = Error> {
    let logger = req.logger();
    let auth = req.request_auth();

    if auth.session.is_none() || auth.account.is_none() {
        return future::result(Ok(KernelError::Unauthorized("Authentication required".to_string()).error_response()))
        .responder();
    }

    return state.db
    .send(controllers::FindSessionsForAccount{
        account_id: auth.account.expect("unwrapping non none account").id,
    })
    .from_err()
    .and_then(move |sessions| {
        match sessions {
            Ok(sessions) => {
                let sessions: Vec<models::Session> = sessions.into_iter().map(|session| {
                    models::Session{
                        id: session.id,
                        created_at: session.created_at,
                        ip: session.ip,
                        location: session.location,
                        device: session.device,
                    }
                }).collect();
                let res = api::Response::data(sessions);
                Ok(HttpResponse::Ok().json(&res))
            },
            Err(err) => Err(err),
        }
    })
    .from_err()
    .map_err(move |err: KernelError| {
        slog_error!(logger, "{}", err);
        return err;
    })
    .from_err()
    .responder();
}
