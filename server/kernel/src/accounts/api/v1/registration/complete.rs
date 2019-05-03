use crate::{
    api,
    accounts::api::v1::models,
    log::macros::*,
    api::middlewares::{
        GetRequestLogger,
        GetRequestId,
    },
    accounts::controllers,
    utils,
};
use std::time::Duration;
use futures::future::Future;
use actix_web::{
    web, Error, HttpRequest, HttpResponse,
};
use futures::future::IntoFuture;
use rand::Rng;


pub fn post(registration_data: web::Json<models::CompleteRegistrationBody>, state: web::Data<api::State>, req: HttpRequest)
-> impl Future<Item = HttpResponse, Error = Error> {
    let mut rng = rand::thread_rng();
    let state = req.state().clone();
    let logger = req.logger();
    let config = state.config.clone();
    let request_id = req.request_id().0;

    // random sleep to prevent bruteforce and sidechannels attacks
    return tokio_timer::sleep(Duration::from_millis(rng.gen_range(400, 650))).into_future()
    .from_err()
    .and_then(move |_|
        state.db
        .send(controllers::CompleteRegistration{
            id: registration_data.id,
            username: registration_data.username.clone(),
            config,
            request_id,
        }).flatten()
    )
    .and_then(|(session, token)| {
        let res = api::Response::data(models::CompleteRegistrationResponse{
            id: session.id,
            token: utils::encode_session(&session.id.to_string(), &token),
        });
        Ok(HttpResponse::Created().json(&res))
    })
    .map_err(move |err| {
        slog_error!(logger, "{}", err);
        return err;
    })
    .from_err()
    .responder();
}
