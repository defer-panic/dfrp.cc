CREATE MIGRATION m1yldhdwiesnmd4bapuza7msgmtvp5b4cqb6hak7jtsbhv6svswjaa
    ONTO m1sdc4oceaujowofmkqg3s5v6olklwpjdaans76tbuh6kp6mpl3jfa
{
  ALTER TYPE default::Shortening {
      ALTER PROPERTY identifier {
          CREATE CONSTRAINT std::exclusive;
      };
  };
};
