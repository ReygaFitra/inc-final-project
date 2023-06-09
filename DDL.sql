CREATE TABLE mst_users(
    user_id serial NOT NULL PRIMARY KEY,
    name character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    password character varying NOT NULL,
    phone_number character varying(20) NOT NULL,
    address character varying(100) NOT NULL,
    balance integer NOT NULL,
    username character varying NOT NULL,
    point integer,
    role character varying,
    CONSTRAINT unique_username UNIQUE(username)
);

CREATE TABLE mst_bank_account(
    account_id serial NOT NULL PRIMARY KEY,
    user_id integer NOT NULL,
    bank_name character varying(50)  NOT NULL,
    account_number character varying(20)  NOT NULL,
    account_holder_name character varying(50) NOT NULL,
    CONSTRAINT bankAcc_userID_fkey FOREIGN KEY(user_id)
    REFERENCES mst_users(user_id)
);

CREATE TABLE mst_card(
    card_id serial NOT NULL PRIMARY KEY,
    user_id integer NOT NULL,
    card_type character varying(10) NOT NULL,
    card_number character varying(20) NOT NULL,
    expiration_date character varying NOT NULL,
    cvv character varying(5) NOT NULL,
    CONSTRAINT card_userID_fkey FOREIGN KEY(user_id)
    REFERENCES mst_users(user_id)
);

CREATE TABLE mst_photo_url(
    photo_id serial NOT NULL PRIMARY KEY,
    url_photo character varying UNIQUE,
    user_id integer,
    CONSTRAINT photo_userID_fkey FOREIGN KEY(user_id)
    REFERENCES mst_users(user_id)
);

CREATE TABLE mst_point_exchange(
    pe_id serial NOT NULL PRIMARY KEY,
    reward character varying(100) NOT NULL,
    price integer NOT NULL
);

CREATE TABLE tx_transaction(
  tx_id serial NOT NULL PRIMARY KEY,
  amount integer,
  transaction_type character varying(20),
  transaction_date date,
  sender_id integer,
  recipient_id integer,
  bank_account_id integer,
  card_id integer,
  pe_id integer,
  point integer,
 
  CONSTRAINT tx_peID_fkey FOREIGN KEY(pe_id)
  REFERENCES mst_point_exchange(pe_id)
);
