//! -*- mode: rust; -*-
//!
//! download - downloads youtube files
use env_logger;
use log::{debug, error, info, log_enabled, Level};
use rustube::{Id, Video};
use std::ffi;

#[no_mangle]
pub async extern "C" fn download<'a>(
    url: &'a str,
    path: &'a str,
) -> Result<(), Box<dyn std::error::Error>> {
    env_logger::init();

    info!("video_url: {:?}", url);
    let id = Id::from_raw(&url)?;
    info!("video_id: {:?}", id);

    let video = Video::from_id(id.into_owned()).await?;
    debug!("raw video: {:?}", video);

    let _result = video
        .streams()
        .iter()
        .filter(|stream| {
            stream.includes_video_track && stream.includes_audio_track
        })
        .max_by_key(|stream| stream.quality_label)
        .unwrap()
        .download_to_dir(&path)
        .await
        .unwrap();
    debug!("download status: {:?}", _result);

    Ok(())
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[tokio::test]
    async fn test_download() {
        let url = String::from("https://www.youtube.com/watch?v=lWEbEtr_Vng");
        let fp = String::from("~/Downloads");
        download(url.as_str(), fp.as_str())
            .await
            .expect("expect an OK(_) response");
    }
}
