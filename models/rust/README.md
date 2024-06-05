# Rust GBFS Language Bindings

![Crates.io Version](https://img.shields.io/crates/v/gbfs_types)
![Crates.io License](https://img.shields.io/crates/l/gbfs_types)

Rust types for parsing and working with General Bikeshare Feed Specification ([GBFS](https://gbfs.org/)) data, ensuring type safety and code consistency in Rust projects.

Most fields required by the GBFS specification are marked as `Option` in this crate, as they may not always present in the data provided by the GBFS feeds.

Some helper functions are provided to fetch the data from the GBFS feeds.

## Installation

To use `gbfs_types` in your own project, you need to install the dependencies

```
cargo add gbfs_types
```

## Versions
Currently only version 3.0 of GBFS is supported

## Example Code (Types)
```rust
use reqwest::Error;
use tokio;
use gbfs_types::v3_0::files::SystemInformationFile;

#[tokio::main]
async fn main() -> Result<(), Error> {
    let gbfs_url = "https://example-gbfs-feed/gbfs/3.0/system-information";
    let client = reqwest::Client::new();
    let response = client.get(url).header("User-Agent", "rust-reqwest").send().await?;

    if response.status().is_success() {
        let system_information: SystemInformationFile = response.json().await?;
        println!("systemInformation version: {}", system_information.version);
        println!("systemInformation data: {}", system_information.data.system_id);
    }
    Ok(())
}
```

## Example Code (Helper Functions)
```rust
use gbfs_types::v3_0::files::gbfs::GbfsDataFeeds;
use gbfs_types::v3_0::files::{GbfsData, SystemInformationFile};
use reqwest::Error;
use tokio;

use gbfs_types::v3_0::urls::SystemInformationFileUrl;
use gbfs_types::v3_0::GbfsObjects;

#[tokio::main]
async fn main() -> Result<(), Error> {
    let gbfs_data: GbfsData = GbfsData {
        feeds: vec![
            GbfsDataFeeds {
                name: "system_information".to_string(),
                url: "https://example-gbfs-feed/gbfs/3.0/system-information".to_string(),
            },
            // ... all other feeds
            GbfsDataFeeds{
                name: "gbfs_versions".to_string(),
                url: "https://example-gbfs-feed/gbfs/3.0/versions".to_string()
            }
        ],
    };
    let system_information = gbfs_data.get_system_information().await;

    match system_information {
        Ok(Some(info)) => {
            // Successfully retrieved SystemInformationFile
            println!("System Information: {:?}", info.data.email);
        }
        Ok(None) => {
            // SystemInformationFile was not present
            println!("No system information available.");
        }
        Err(e) => {
            // Handle the error
            eprintln!("Error retrieving system information: {}", e);
        }
    }
    Ok(())
}

```

## Contributions
We extend our gratitude to [Fluctuo](https://fluctuo.com) for generously providing us with this repository. Their support and contribution have been instrumental in advancing our project and fostering collaboration within the community. Thank you, Fluctuo, for your dedication to open source and for empowering developers worldwide.