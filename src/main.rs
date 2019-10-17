mod app;
mod info;
mod jobs;
mod server;

use bloom_kernel::{
    config,
    log,
    // myaccount::domain::account,
};
use clap::{App as ClapApp, Arg, SubCommand};
use sentry::integrations::panic::register_panic_handler;
use std::env;
use std::error::Error;

fn register_reactors() {
    // eventsourcing::subscribe::<_, account::Created, _>(Box::new(
    //     drive::reactors::AccountCreated {},
    // ));
    // eventsourcing::subscribe::<_, account::Created, _>(Box::new(
    //     bitflow::reactors::AccountCreated {},
    // ));
    // eventsourcing::subscribe::<_, account::Event, _>(Box::new(
    //     billing::reactors::AccountCreated {},
    // ));
}

fn main() -> Result<(), Box<dyn Error>> {
    // init config
    let cfg = config::init();

    // init sentry (error tracking)
    let _sentry_guard = sentry::init(cfg.sentry.server_url.clone());
    env::set_var("RUST_BACKTRACE", "1");
    register_panic_handler();

    // init other parts
    let (_log_guard, logger) = log::setup_slog();
    crypto42::init()?;
    register_reactors();

    let matches = ClapApp::new(info::NAME)
        .version(info::VERSION)
        .author(info::AUTHOR)
        .about(info::ABOUT)
        .subcommand(SubCommand::with_name("server").about("run the server"))
        .subcommand(
            SubCommand::with_name("jobs")
                .about("list all jobs")
                .subcommand(SubCommand::with_name("ls").about("list all jobs"))
                .subcommand(
                    SubCommand::with_name("run")
                        .about("run the specific job")
                        .arg(Arg::with_name("job_name").required(true)),
                ),
        )
        .get_matches();

    if matches.subcommand_matches("server").is_some() {
        return server::run(cfg, logger);
    } else if let Some(matches) = matches.subcommand_matches("jobs") {
        if matches.subcommand_matches("ls").is_some() {
            return jobs::list_jobs();
        } else if let Some(matches) = matches.subcommand_matches("run") {
            return jobs::run_job(
                cfg,
                &matches
                    .value_of("job_name")
                    .expect("error getting job_name argument"),
            );
        } else {
            println!(
                "{}",
                matches
                    .usage
                    .as_ref()
                    .expect("error getting jobs command usage")
            );
        }
    }

    return Ok(());
}
