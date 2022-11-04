CREATE MIGRATION m1sdc4oceaujowofmkqg3s5v6olklwpjdaans76tbuh6kp6mpl3jfa
    ONTO m1mwqntei2nia7fzqo46ky6jw46zh56rmblwdvfiyyyetuj6nisuoa
{
  ALTER TYPE default::Shortening {
      CREATE INDEX ON (.identifier);
  };
};
