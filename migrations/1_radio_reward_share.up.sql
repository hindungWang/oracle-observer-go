CREATE TABLE radio_reward_shares (
    amount bigInt NOT NULL,
    start_epoch bigInt NOT NULL,
    end_epoch bigInt NOT NULL,
    hotspot_key BYTEA NOT NULL,
    cbsd_id varchar(52) NOT NULL,
    PRIMARY KEY (end_epoch, hotspot_key, cbsd_id)
);
