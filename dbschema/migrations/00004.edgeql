CREATE MIGRATION m1hao56ls7a26c2rlmj2dh4a6fwsdavqvdfty6bpois4bfteybgfsq
    ONTO m1yldhdwiesnmd4bapuza7msgmtvp5b4cqb6hak7jtsbhv6svswjaa
{
  ALTER TYPE default::Shortening {
      CREATE PROPERTY updated_at -> std::datetime;
      CREATE REQUIRED PROPERTY visits -> std::int64 {
          SET REQUIRED USING (0);
      };
  };
};
