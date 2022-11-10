CREATE MIGRATION m1dtuzuob3ly5qbwxafpfql2fa67ea6ms6jvzoydrxl5d6gzxvi36q
    ONTO m1mdgtqzcteikoqcuxydqiuq6s5gifc3n6wcekow75t4yw2c2bxala
{
  ALTER TYPE default::Shortening {
      ALTER PROPERTY visits {
          SET default := 0;
          SET REQUIRED USING (0);
      };
  };
};
