CREATE MIGRATION m1mdgtqzcteikoqcuxydqiuq6s5gifc3n6wcekow75t4yw2c2bxala
    ONTO m12yqd42iwut3pkzvv332owy5rextfugiotd2v3paixsk6jdeqwmqq
{
  CREATE TYPE default::User {
      CREATE REQUIRED PROPERTY created_at -> std::datetime {
          SET default := (std::datetime_of_transaction());
      };
      CREATE PROPERTY gh_access_key -> std::str;
      CREATE REQUIRED PROPERTY gh_login -> std::str {
          CREATE CONSTRAINT std::exclusive;
      };
      CREATE REQUIRED PROPERTY is_active -> std::bool {
          SET default := true;
      };
  };
  ALTER TYPE default::Shortening {
      CREATE LINK created_by -> default::User;
      ALTER PROPERTY created_at {
          SET default := (std::datetime_of_transaction());
      };
  };
};
