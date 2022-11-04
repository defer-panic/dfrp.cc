CREATE MIGRATION m1mwqntei2nia7fzqo46ky6jw46zh56rmblwdvfiyyyetuj6nisuoa
    ONTO initial
{
  CREATE TYPE default::Shortening {
      CREATE REQUIRED PROPERTY created_at -> std::datetime;
      CREATE REQUIRED PROPERTY identifier -> std::str;
      CREATE REQUIRED PROPERTY original_url -> std::str;
  };
};
