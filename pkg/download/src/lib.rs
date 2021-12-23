//! -*- mode: rust; -*-
//!
//! download - downloads youtube files
use rustube::{Id, Video};
use std::ffi;

#[no_mangle]
pub async extern "C" fn download<'a>(
    url: &'a str,
    path: &'a str,
) -> Result<(), Box<dyn std::error::Error>> {
    env_logger::init();

    let id = Id::from_raw(&url)?;
    let video = Video::from_id(id.into_owned()).await?;

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

    Ok(())
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[tokio::test]
    async fn test_download() {
        let url = String::from("https://www.youtube.com/watch?v=lWEbEtr_Vng");
        let fp = String::from("~/Downloads");
        download(url.as_str(), fp.as_str()).await.expect("expect an OK(_) response");
    }
}
